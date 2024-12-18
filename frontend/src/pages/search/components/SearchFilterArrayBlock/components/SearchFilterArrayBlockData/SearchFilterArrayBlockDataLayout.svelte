<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps= $props();
</script>

<div class="w-svw sm:w-112 lg:w-search-sidebar flex flex-col">
    {@render children()}
</div>