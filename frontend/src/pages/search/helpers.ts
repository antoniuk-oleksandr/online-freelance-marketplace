import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import { searchFilterDrawerStore } from "@/pages/search/stores/search-filter-drawer-store";
import { searchStore } from "@/pages/search/stores/search-store";
import { Category } from "@/types/Category";
import { Language } from "@/types/Language";
import { OrderType } from "@/types/OrderType";
import type { SearchFilterArrayInputDataStore } from "@/types/SearchFilterArrayInputDataStore";
import type { SearchFilterArrayStore } from "@/types/SearchFilterArrayStore";
import { SearchPageParams } from "@/types/SearchPageParams";
import type { SearchRequestResponse } from "@/types/SearchRequestResponse";
import type { SelectItem } from "@/types/SelectItem";
import { Skill } from "@/types/Skill";
import { SortType } from "@/types/SortType";
import { onDestroy } from "svelte";
import { navigate } from "svelte-routing";

export const getSearchPageParams = (): SearchPageParams => {
    const params = new URLSearchParams(window.location.search);

    return {
        query: params.get("query"),
        sort: params.get("sort") ? parseInt(params.get("sort")!) as SortType : null,
        order: params.get("order") ? parseInt(params.get("order")!) as OrderType : null,
        skill: params.getAll("skill"),
        priceFrom: params.get("priceFrom"),
        priceTo: params.get("priceTo"),
        levelFrom: params.get("levelFrom"),
        levelTo: params.get("levelTo"),
        ratingFrom: params.get("ratingFrom"),
        ratingTo: params.get("ratingTo"),
        deliveryTimeFrom: params.get("deliveryTimeFrom"),
        deliveryTimeTo: params.get("deliveryTimeTo"),
        category: params.getAll("category"),
        language: params.getAll("language"),
    }
}

export const useIsClosedAfterAnimation = (
    breakpoint: number,
    setIsMobile: (value: boolean) => void,
) => {
    const handleChange = (_: Event) => {
        setIsMobile(window.innerWidth < breakpoint)
    }

    window.addEventListener("resize", handleChange);

    onDestroy(() => window.removeEventListener("resize", handleChange));
}

export const makeSearchSelectTitle = (item: number, isOrder?: boolean) => {
    return "By " + Object.values(isOrder ? OrderType : SortType)[item].toLowerCase();
}

export const makeASearchSelectItemList = (
    searchPageParams: SearchPageParams | undefined,
    isOrder?: boolean
): SelectItem[] => {
    if (!searchPageParams) return [];

    return Array
        .from({ length: Object.keys(isOrder ? OrderType : SortType).length / 2 })
        .map((_, index) => ({
            title: makeSearchSelectTitle(index, isOrder),
            clickAction: () =>
                setSelectClickAction(searchPageParams, index, isOrder)
        }));
}

export const setSelectClickAction = (
    searchPageParams: SearchPageParams,
    index: number,
    isOrder?: boolean,
) => {
    if (isOrder) searchPageParams.order = index;
    else searchPageParams.sort = index;

    submitSearchPage(searchPageParams);
}

export const submitSearchPage = (
    searchPageParams?: SearchPageParams
) => {
    if (!searchPageParams) return;

    const newLink = generateSearchLink(searchPageParams);
    navigate(newLink);
}

export const resetSearchPage = () => {
    navigate("/search");
    searchFilterDrawerStore.set(false);
}

const appendSearchLinkParam = (
    searchPageParams: SearchPageParams,
    attribute: keyof SearchPageParams,
    params: URLSearchParams
) => {
    const value = searchPageParams[attribute];

    if (Array.isArray(value)) {
        value.forEach((item) => {
            params.append(attribute, item.toString());
        })
    } else if (value) params.append(attribute, value.toString());
}

export const generateSearchLink = (
    searchPageParams: SearchPageParams,
    cursor?: string | null,
) => {
    const params = new URLSearchParams();

    const attributes = ["query", "page", "sort", "order", "skill", "language", "category", "priceFrom", "priceTo", "levelFrom", "levelTo", "ratingFrom", "ratingTo", "deliveryTimeFrom", "deliveryTimeTo", "cursor"];

    attributes.forEach((item) => {
        appendSearchLinkParam(searchPageParams, item as keyof SearchPageParams, params);
    })

    return `/search${cursor ? `?cursor=${cursor}&` : "?"}${params.toString()}`;
}

export const getFormFromInputValue = (
    searchPageParams: SearchPageParams,
    item: string,
    type: "From" | "To",
) => {
    const attribute = getFormToInputAttribute(item, type);

    return parseInt(searchPageParams[attribute] as string);
}

export const setFormFromInputValue = (
    value: number,
    item: string,
    type: "From" | "To",
) => {
    const attribute = getFormToInputAttribute(item, type);

    searchStore.update((store) => {
        if (!store) return store;
        else return ({
            ...store,
            [attribute]: value,
        })
    });
}

export const getFormToInputAttribute = (
    item: string,
    type: "From" | "To",
): keyof SearchPageParams => {
    return `${item}${type}` as keyof SearchPageParams;
}

export const getSelectedSearchSelectItem = (
    searchPageParams: SearchPageParams | undefined,
    isOrder?: boolean,
) => {
    if (!searchPageParams) return "";

    const searchAttribute = isOrder ? "order" : "sort";
    const value = searchPageParams[searchAttribute] ? searchPageParams[searchAttribute] : 0;

    const val = Object.values(isOrder ? OrderType : SortType)[value];

    return "By " + val?.toString().toLowerCase();
}

export const resetSearchArrayParam = (
    storeData: SearchFilterArrayStore,
) => {
    searchStore.update((prev) => {
        if (!prev) return prev;

        let copy = { ...prev };
        copy[storeData.attribute] = [];

        return copy;
    })
}

export const getAllFilterArrayBlockDataElements = (
    defaultFilterParams: { skill: Skill[], category: Category[], language: Language[] },
    storeData: SearchFilterArrayStore,
    inputValues: SearchFilterArrayInputDataStore | undefined,
) => {
    if (!inputValues || !defaultFilterParams) return undefined;

    return defaultFilterParams[storeData.attribute].filter((item) =>
        item.name.toLowerCase()
            .includes(inputValues[storeData.attribute].toLowerCase())
    )
}

export const makeSearchRequest = (
    searchPAgeParams: SearchPageParams,
    setSearchRequestResponse: (value: SearchRequestResponse) => void,
) => {
    const link = generateSearchLink(searchPAgeParams);
    request<SearchRequestResponse>("GET", link).then((response) => {
        if (response.status !== 200) {
            errorStore.set({ shown: true, error: response.data.error });
        }

        setSearchRequestResponse(response);
    });
}
