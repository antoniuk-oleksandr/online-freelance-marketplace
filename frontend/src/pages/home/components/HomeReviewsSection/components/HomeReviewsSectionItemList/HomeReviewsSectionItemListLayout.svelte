<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
    {@render children()}
</div>