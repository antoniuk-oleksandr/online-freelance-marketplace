<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="grid grid-cols-1 md:grid-cols-3 gap-12">
    {@render children()}
</div>