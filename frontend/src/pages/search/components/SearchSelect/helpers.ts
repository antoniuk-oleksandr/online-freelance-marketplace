import type { NewSelectItem } from "@/types/NewSelectItem";
import { OrderType } from "@/types/OrderType";
import type { SearchPageParams } from "@/types/SearchPageParams";
import { SortType } from "@/types/SortType";
import { submitSearchPage } from "../../helpers";
import { capitalize } from "@/utils/utils";

export const getSearchSelectItems = (): NewSelectItem[] => {
  const sortSelectItems = makeSortSelectItems();
  const orderSelectItems = makeOrderSelectItems();

  orderSelectItems[0].dividerBefore = true;

  return [...sortSelectItems, ...orderSelectItems];
}

const makeSortSelectItems = () => Object.values(SortType).reduce((acc, key) => {
  if (Number.isNaN(parseInt(key as string))) {
    acc.push({
      title: `By ${key.toLocaleString().toLowerCase()}`,

    });
  }
  return acc;
}, [] as NewSelectItem[])

const makeOrderSelectItems = () => Object.values(OrderType).reduce((acc, key) => {
  if (Number.isNaN(parseInt(key as string))) {
    acc.push({
      title: `By ${key.toLocaleString().toLowerCase()}`,
      groupId: 1,
    });
  }
  return acc;
}, [] as NewSelectItem[])


export const handleSearchSelectChange = (
  value: string | string[],
  searchPageParams: SearchPageParams | undefined,
) => {
  if (!Array.isArray(value) || !searchPageParams) return;

  value = value.map((item) => capitalize(item.replaceAll('By ', '')));

  searchPageParams.sort = convertStringToSortType(value[0]);
  searchPageParams.order = convertStringToOrderType(value[1]);

  submitSearchPage(searchPageParams);
}

const convertStringToSortType = (value: string): SortType =>
  Object.entries(SortType).reduce((acc, [key, val]) => {
    if (value === key) acc = val as SortType;
    return acc;
  }, 0 as SortType)

const convertStringToOrderType = (value: string): OrderType =>
  Object.entries(OrderType).reduce((acc, [key, val]) => {
    if (value === key) acc = val as OrderType;
    return acc;
  }, 0 as OrderType)

export const makeInitialSearchSelectIndexes = (
  searchPageParams: SearchPageParams | undefined,
): number[] => {
  if (searchPageParams === undefined) return [0, 5];
  const order = upscaleOrderVal(searchPageParams.order);
  const sort = getSortInitialVal(searchPageParams.sort) ;

  return [sort as number, order];
}

const getSortInitialVal = (value: number | undefined | null) => {
  if(value === undefined || value === null) return SortType.Popularity;
  else return value;
}

const upscaleOrderVal = (val: number | undefined | null) => {
  if (val === undefined || val === null) return Object.keys(SortType).length / 2;
  else return val + Object.keys(SortType).length / 2;
}

export const makeSearchSelectTriggerIcon = (
  searchPageParams: SearchPageParams | undefined,
) => {
  if(!searchPageParams) return 'hugeicons:sort-by-up-01';
  if(searchPageParams.order === null || searchPageParams.order === OrderType.Descending) return 'hugeicons:sort-by-down-01';
  if(searchPageParams.order === OrderType.Ascending) return 'hugeicons:sort-by-up-01';
}
