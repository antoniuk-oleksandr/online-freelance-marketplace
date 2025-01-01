import {writable} from "svelte/store";
import {UserByIdData} from "@/types/GetUserByIdData.ts";

export const userByIdStore = writable<UserByIdData | undefined>();