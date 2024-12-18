<script lang="ts">
    import PageLayout from "@/common-components/PageLayout.svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<footer
        class={"border-t py-4 md:py-0 bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider"}
>
    <PageLayout>
        <div class={"flex  md:h-16 md:flex-row w-full flex-col gap-y-4 justify-between items-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            {@render children()}
        </div>
    </PageLayout>
</footer>