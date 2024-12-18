<script lang="ts">
    import PageLayout from "@/common-components/PageLayout.svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<header class={"h-16 sticky top-0 z-30 flex items-center bg-light-palette-background-block dark:bg-dark-palette-background-block border-b border-light-palette-divider dark:border-dark-palette-divider shadow-sm w-full"}>
    <PageLayout>
        <div class={"flex w-full items-center justify-between"}>
            {@render children()}
        </div>
    </PageLayout>
</header>