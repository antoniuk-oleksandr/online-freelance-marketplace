import { writable } from "svelte/store";

export const chatBodyStore = writable<HTMLDivElement | null>(null)
