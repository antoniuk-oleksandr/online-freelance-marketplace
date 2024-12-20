import {writable} from "svelte/store";
import {SearchCursorStore} from "@/types/SearchCursorStore.ts";

export const searchCursorStore = writable<SearchCursorStore>({
    originalSearchPageParams: undefined,
    searchRequestResponse: undefined,
});