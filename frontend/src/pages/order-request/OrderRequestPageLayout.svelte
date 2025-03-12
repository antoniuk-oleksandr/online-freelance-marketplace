<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex flex-col gap-6 text-base h-full">
    {@render children()}
</div>
