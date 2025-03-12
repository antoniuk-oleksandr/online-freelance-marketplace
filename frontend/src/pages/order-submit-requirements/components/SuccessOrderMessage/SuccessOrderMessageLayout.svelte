<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="bg-green-400 flex gap-6 items-center bg-opacity-30 dark:bg-green-400 dark:bg-opacity-30 p-6 rounded-xl">
    {@render children()}
</div>
