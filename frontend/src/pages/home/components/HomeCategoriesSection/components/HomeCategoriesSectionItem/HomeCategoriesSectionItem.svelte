<script lang="ts">
    import Icon from "@iconify/svelte";
    import CategoriesSectionItemLayout
        from "@/pages/home/components/HomeCategoriesSection/components/HomeCategoriesSectionItem/HomeCategoriesSectionItemLayout.svelte";

    type CategoriesSectionItem = {
        icon: string,
        title: string,
        id: number
    }

    const {icon, title, id}: CategoriesSectionItem = $props();
</script>

<CategoriesSectionItemLayout id={id}>
    <Icon icon={icon} class="text-cyan-500 dark:text-cyan-500" width="32" height="32"/>
    <h3
            class="text-sm font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary"
    >{title}</h3>
</CategoriesSectionItemLayout>