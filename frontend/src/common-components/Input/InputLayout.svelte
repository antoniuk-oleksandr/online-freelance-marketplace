<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} :LayoutProps = $props();
</script>

<div class="flex gap-y-1 flex-col">
    {@render children()}
</div>