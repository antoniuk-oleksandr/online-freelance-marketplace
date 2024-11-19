<script lang="ts">
    import type {Package} from "@/types/Package.ts";
    import PackageTabListLayout from "@/pages/services/components/PackageTabList/PackageTabListLayout.svelte";
    import PackageTab from "@/pages/services/components/PackageTab/PackageTab.svelte";

    type PackageTabListProps = {
        packages: Package[],
        tabs: string[],
        selectedTab: number,
        setSelectedTab: (value: number) => void,
    }

    const listProps: PackageTabListProps = $props();
    const {packages, tabs}: PackageTabListProps = listProps;
</script>

<PackageTabListLayout>
    {#each packages as _, index}
        <PackageTab
                {...listProps}
                index={index}
                length={packages.length}
                text={tabs[index]}
        />
    {/each}
</PackageTabListLayout>
