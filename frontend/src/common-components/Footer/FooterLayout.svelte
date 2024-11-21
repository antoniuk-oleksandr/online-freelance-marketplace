<script>
    import PageLayout from "@/common-components/PageLayout.svelte";
</script>

<footer
        class={"border-t bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider py-4"}
>
    <PageLayout>
        <div class={"flex md:flex-row w-full flex-col gap-y-4 justify-between items-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            <slot/>
        </div>
    </PageLayout>
</footer>