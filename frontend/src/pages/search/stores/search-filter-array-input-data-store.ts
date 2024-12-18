import {writable} from "svelte/store";
import {SearchFilterArrayInputDataStore} from "@/types/SearchFilterArrayInputDataStore.ts";

export const searchFilterArrayInputDataStore = writable<SearchFilterArrayInputDataStore>({
    skill: "",
    language: "",
    category: ""
})