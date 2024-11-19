<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="aspect-video h-9 w-auto">
    {@render children()}
</div>
