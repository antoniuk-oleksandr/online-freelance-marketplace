<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="mt-10 flex justify-between items-end">
    {@render children()}
</div>