import type { SignHeaderData } from "@/types/SignHeaderData";
import { writable } from "svelte/store";
import { getUserSession } from "../../helpers";

export const signDataStore = writable<SignHeaderData | undefined>()
