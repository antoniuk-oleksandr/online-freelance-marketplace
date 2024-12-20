<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {searchCursorStore} from "@/pages/search/stores/search-cursor-store.ts";

    const {children}: LayoutProps = $props();
    let isScrollPresent = $state<boolean | undefined>();

    searchCursorStore.subscribe((value) => {
        if(!value.searchRequestResponse) return;

        const {scrollHeight, clientHeight} = document.documentElement;
        isScrollPresent = scrollHeight > clientHeight;
    })
</script>

<div class="flex min-h-search-left-side-mobile w-full {isScrollPresent ? 'mb-6' : ''}">
    {@render children()}
</div>