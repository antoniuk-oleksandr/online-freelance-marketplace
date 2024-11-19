<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();
</script>

<div class="flex gap-4 items-center">
    {@render children()}
</div>