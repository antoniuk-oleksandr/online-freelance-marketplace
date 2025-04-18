import type { SignHeaderData } from "@/types/SignHeaderData";
import { writable } from "svelte/store";

export const signDataStore = writable<SignHeaderData | undefined>()
