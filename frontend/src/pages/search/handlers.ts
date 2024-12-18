import {searchStore} from "@/pages/search/stores/search-store.ts";
import {SearchPageParams} from "@/types/SearchPageParams.ts";
import {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
import {searchFilterArrayStore} from "@/pages/search/stores/search-filter-array-store.ts";
import {Language} from "@/types/Language.ts";
import {Skill} from "@/types/Skill.ts";
import {Category} from "@/types/Category.ts";
import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
import {submitSearchPage} from "@/pages/search/helpers.ts";
import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";
import {searchFilterArrayInputDataStore} from "@/pages/search/stores/search-filter-array-input-data-store.ts";

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

export const handleSearchScroll = (
    reached: boolean,
    setReached: (value: boolean) => void,
) => {
    if (reached) return;

    const scrollTop = document.documentElement.scrollTop;
    const scrollHeight = document.documentElement.scrollHeight;
    const clientHeight = document.documentElement.clientHeight;

    const scrolledPercent = (scrollTop + clientHeight) / scrollHeight;

    const isPast75Percent = scrolledPercent >= 0.75;

    if (isPast75Percent) {
        console.log("You have scrolled past 75%!");
        setReached(true);
    }
}