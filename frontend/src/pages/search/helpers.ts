import {SearchPageParams} from "@/types/SearchPageParams.ts";
import {SortType} from "@/types/SortType.ts";
import {OrderType} from "@/types/OrderType.ts";
import {onDestroy} from "svelte";
import type {DropdownItem} from "@/types/DropdownItem.ts";
import {navigate} from "svelte-routing";
import {searchStore} from "@/pages/search/stores/search-store.ts";
import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
import {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
import type {SearchFilterArrayInputDataStore} from "@/types/SearchFilterArrayInputDataStore.ts";
import {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";

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

export const makeSearchDropdownTitle = (item: number, isOrder?: boolean) => {
    return "By " + Object.values(isOrder ? OrderType : SortType)[item].toLowerCase();
}

export const makeASearchDropdownItemList = (
    searchPageParams: SearchPageParams | undefined,
    isOrder?: boolean
): DropdownItem[] => {
    if (!searchPageParams) return [];

    return Array
        .from({length: Object.keys(isOrder ? OrderType : SortType).length / 2})
        .map((_, index) => ({
            title: makeSearchDropdownTitle(index, isOrder),
            clickAction: () =>
                setDropDownClickAction(searchPageParams, index, isOrder)
        }));
}

export const setDropDownClickAction = (
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

const generateSearchLink = (
    searchPageParams: SearchPageParams
) => {
    const params = new URLSearchParams();

    const attributes = ["query", "page", "sort", "order", "skill", "language", "category", "priceFrom", "priceTo", "levelFrom", "levelTo", "ratingFrom", "ratingTo", "deliveryTimeFrom", "deliveryTimeTo"];

    attributes.forEach((item) => {
        appendSearchLinkParam(searchPageParams, item as keyof SearchPageParams, params);
    })

    return "/search?" + params.toString();
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

export const getSelectedSearchDropdownItem = (
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

        let copy = {...prev};
        copy[storeData.attribute] = [];

        return copy;
    })
}

export const getAllFilterArrayBlockDataElements = (
    defaultFilterParams: GetFilterParamsRequestResponse["data"],
    storeData: SearchFilterArrayStore,
    inputValues: SearchFilterArrayInputDataStore | undefined,
) => {
    if (!inputValues || !defaultFilterParams) return undefined;

    return defaultFilterParams[storeData.attribute].filter((item) =>
        item.name.toLowerCase()
            .includes(inputValues[storeData.attribute].toLowerCase())
    )
}