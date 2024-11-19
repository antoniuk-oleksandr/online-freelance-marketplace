<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex flex-col gap-y-2 text-base">
    {@render children()}
</div>