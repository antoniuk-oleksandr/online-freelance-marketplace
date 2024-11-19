<script lang="ts">
    import {onMount} from "svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();

    let closed = $state(true);
    let maxHeight = $state(0);
    let minHeight = $state<number | null>(null);
    let content: HTMLElement | null = $state(null);

    onMount(() => {
        if (!content) return;


        maxHeight = content.scrollHeight;
        if (content.scrollHeight > 72) minHeight = 72;
        else minHeight = content.scrollHeight;

    });

    function toggle() {
        closed = !closed;
    }
</script>

<div>
    <div
            bind:this={content}
            class="overflow-hidden transition-all duration-1000 ease-in-out"
            style="height: {closed ? `${minHeight}px` : `${maxHeight}px`}"
    >
        {@render children()}
    </div>

    {#if minHeight === 72}
        <div class="flex justify-center mt-2">
            <button
                    class="text-cyan-500 hover:underline duration-200"
                    onclick={toggle}
            >
                {closed ? "Show more" : "Hide"}
            </button>
        </div>
    {/if}
</div>
