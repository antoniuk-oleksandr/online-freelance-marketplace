import type { ChatMessage } from "@/types/ChatMessage";
import { writable } from "svelte/store";

export const messageStore = writable<ChatMessage[]>([])
