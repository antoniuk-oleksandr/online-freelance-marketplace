import { searchStore } from "@/pages/search/stores/search-store";
import { SearchPageParams } from "@/types/SearchPageParams";
import { SearchFilterArrayAttribute } from "@/types/SearchFilterArrayAttribute";
import { searchFilterArrayStore } from "@/pages/search/stores/search-filter-array-store";
import { Language } from "@/types/Language";
import { Skill } from "@/types/Skill";
import { Category } from "@/types/Category";
import { generateSearchLink, getSearchPageParams, makeSearchRequest, submitSearchPage } from "@/pages/search/helpers";
import { searchFilterDrawerStore } from "@/pages/search/stores/search-filter-drawer-store";
import { searchFilterArrayInputDataStore } from "@/pages/search/stores/search-filter-array-input-data-store";
import type { SearchRequestResponse } from "@/types/SearchRequestResponse";
import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import { searchCursorStore } from "@/pages/search/stores/search-cursor-store";
import { SearchCursorStore } from "@/types/SearchCursorStore";

export const handleSearchFromToBlockInputSpinnersClick = (
    value: number,
    setValue: (value: number) => void,
    direction: "up" | "down"
) => {
    if (direction === "down") {
        if (Number.isNaN(value)) setValue(0);
        else if (value - 1 >= 0) setValue(value - 1)
    } else if (direction === "up") {
        if (Number.isNaN(value)) setValue(1);
        else setValue(value + 1)
    }
}

export const handleSearchFilterArrayBlockClick = (
    data: Language[] | Skill[] | Category[],
    title: string,
    attribute: SearchFilterArrayAttribute,
) => {
    searchFilterArrayStore.update(() => ({
        data,
        title,
        attribute,
        isOpened: true
    }))
}

export const handleSearchFilterArrayBlockDataElementClick = (
    id: string,
    attribute: SearchFilterArrayAttribute
) => {
    searchStore.update((prev) => {
        if (!prev) return prev;

        if (prev[attribute].includes(id)) {
            return {
                ...prev,
                [attribute]: prev[attribute].filter((item) => item !== id)
            }
        } else return {
            ...prev,
            [attribute]: [...prev[attribute], id]
        }
    });
}

export const handleSearchArrSelectButtonClick = (
    searchPageParams?: SearchPageParams
) => {
    submitSearchPage(searchPageParams);

    searchFilterArrayStore.update((prev) => ({
        ...prev,
        isOpened: false
    }))

    setTimeout(() => {
        searchFilterDrawerStore.set(false);
    }, 350)
}

export const handleSearchFilterArrayInput = (
    attribute: SearchFilterArrayAttribute,
    e?: Event & { currentTarget: (EventTarget & HTMLInputElement) },
) => {
    if (!e) return;

    searchFilterArrayInputDataStore.update((prev) => {
        return {
            ...prev,
            [attribute]: e.currentTarget.value
        }
    })
}

export const handleSearchInput = (
    e: Event & { currentTarget: (EventTarget & HTMLInputElement) }
) => {
    const value = e.currentTarget.value;

    searchStore.update((prev) => {
        if (!prev) return prev;

        return {
            ...prev,
            query: value
        }
    });
}

export const createLastServiceObserver = (
    gridElement: HTMLElement | undefined,
    searchCursorData: SearchCursorStore,
) => {
    if (!gridElement) return;

    const lastChildElement = gridElement.children[gridElement.childElementCount - 1]

    const observer = new IntersectionObserver((entries) => {
        const last = entries[0];
        if (!last.isIntersecting || !searchCursorData.originalSearchPageParams || !searchCursorData.searchRequestResponse) return;
        let { searchRequestResponse, originalSearchPageParams } = searchCursorData;

        if (searchRequestResponse.status === 200 && searchRequestResponse.data.hasMore) {
            const link = generateSearchLink(originalSearchPageParams, searchRequestResponse.data.cursor);
            request<SearchRequestResponse>('GET', link).then((response) => {
                if (response.status !== 200) {
                    errorStore.set({ shown: true, error: response.data.error });
                }

                if (response.status !== 200 || !searchRequestResponse.data.services || !response.data.services) return;

                response.data.services = [...searchRequestResponse.data.services, ...response.data.services]

                searchCursorStore.update((prev) => ({
                    ...prev,
                    searchRequestResponse: response,
                    cursorAdded: true
                }))
            });
        }

        observer.unobserve(last.target)
    }, {
        rootMargin: "1600px"
    })

    observer.observe(lastChildElement);
}

export const handleSearchRefresh = () => {
    const params = getSearchPageParams();
    searchStore.set(params);

    searchCursorStore.update((prev) => ({
        ...prev,
        cursorAdded: false,
        originalSearchPageParams: params
    }));


    const link = generateSearchLink(params);
    processRequest(link);
}

const processRequest = (link: string) => {
    request<SearchRequestResponse>("GET", link).then((response) => {
        if (response.status === 200) {
            searchCursorStore.update((prev) => {
                if (!prev.originalSearchPageParams) return prev;

                return {
                    ...prev,
                    cursorAdded: true,
                    searchRequestResponse: response
                }
            })
        } else {
            errorStore.set({ shown: true, error: response.data.error });
        }
    });
}
