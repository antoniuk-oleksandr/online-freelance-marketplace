<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex-shrink-0 w-full">
    {@render children()}
</div>