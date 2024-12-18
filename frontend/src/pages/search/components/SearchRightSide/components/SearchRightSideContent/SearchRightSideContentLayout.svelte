<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex flex-col h-full w-svw sm:w-112 lg:w-search-sidebar">
    {@render children()}
</div>