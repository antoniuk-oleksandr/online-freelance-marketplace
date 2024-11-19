<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="md:w-20  flex items-center flex-row-reverse md:block">
    {@render children()}
</div>