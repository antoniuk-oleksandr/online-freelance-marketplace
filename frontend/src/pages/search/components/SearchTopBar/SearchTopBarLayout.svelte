<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="items-center justify-between grid grid-cols-search-top-bar gap-3">
    {@render children()}
</div>