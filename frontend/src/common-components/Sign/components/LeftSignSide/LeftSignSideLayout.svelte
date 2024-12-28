<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex flex-col justify-center gap-y-6 px-4 md:px-24">
    {@render children()}
</div>