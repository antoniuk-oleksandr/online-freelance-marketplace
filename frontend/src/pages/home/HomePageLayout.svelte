<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex gap-y-6 w-full flex-col text-base">
    {@render children()}
</div>