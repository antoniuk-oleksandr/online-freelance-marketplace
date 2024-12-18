import { writable } from "svelte/store";
import {SearchPageParams} from "@/types/SearchPageParams.ts";

export const searchStore = writable<SearchPageParams | undefined>();