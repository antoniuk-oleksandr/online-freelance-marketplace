<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="text-base flex flex-col gap-y-2 py-2 px-6">
    {@render children()}
</div>