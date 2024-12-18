<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type SearchLeftSideLayoutProps = LayoutProps & {
        searchRequestResponse: any;
    }

    const {searchRequestResponse, children}: SearchLeftSideLayoutProps = $props();
    let isScrollPresent = $state<boolean | undefined>();

    $effect(() => {
        if (searchRequestResponse) {
            const {scrollHeight, clientHeight} = document.documentElement;
            isScrollPresent = scrollHeight > clientHeight;
        }
    });
</script>

<div class="flex min-h-search-left-side-mobile w-full {isScrollPresent ? 'mb-6' : ''}">
    {@render children()}
</div>