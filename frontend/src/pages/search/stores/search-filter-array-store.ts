import {writable} from "svelte/store";
import {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";

export const searchFilterArrayStore = writable<SearchFilterArrayStore>({
    data: [],
    title: "",
    isOpened: false,
    attribute: "skill",
});