<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {createLastServiceObserver} from "@/pages/search/handlers.ts";
    import type {SearchCursorStore} from "@/types/SearchCursorStore.ts";
    import {searchCursorStore} from "@/pages/search/stores/search-cursor-store.ts";

    const {children}: LayoutProps = $props();

    let gridElement = $state<HTMLElement | undefined>();

    let searchCursorData = $state<SearchCursorStore | undefined>();
    searchCursorStore.subscribe((value) => searchCursorData = value);

    $effect(() => {
        if(!searchCursorData) return;
        createLastServiceObserver(gridElement, searchCursorData);
    })
</script>

<div
        bind:this={gridElement}
        id="services-grid"
        class="grid grid-cols-1 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-3 xl:grid-cols-3 2xl:grid-cols-4 gap-3"
>
    {@render children()}
</div>