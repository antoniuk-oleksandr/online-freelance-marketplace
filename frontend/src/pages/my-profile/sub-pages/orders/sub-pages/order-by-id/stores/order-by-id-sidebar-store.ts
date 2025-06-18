import type { OrderByIdSidebarStore } from "@/types/OrderByIdSidebarStore";
import { writable } from "svelte/store";

export const orderByIdSidebarStore = writable<OrderByIdSidebarStore | undefined>();
