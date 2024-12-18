<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex flex-col overflow-scroll max-h-search-array-mobile lg:max-h-search-array">
    {@render children()}
</div>