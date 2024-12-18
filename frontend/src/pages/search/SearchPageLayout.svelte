<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children }: LayoutProps = $props();
</script>

<div class="grid grid-cols-1 h-full lg:grid-cols-search-page gap-6 flex-grow">
    {@render children()}
</div>