<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex flex-col-reverse lg:flex-col lg:grid grid-cols-search-page gap-6 h-full">
    {@render children()}
</div>
