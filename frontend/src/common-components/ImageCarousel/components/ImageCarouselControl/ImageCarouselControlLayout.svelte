<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="absolute top-1/2 left-0 right-0 flex justify-between transform -translate-y-1/2 px-4">
    {@render children()}
</div>