<script lang="ts">
    import ServicePackagesBlockLayout
        from "@/pages/services/components/ServicePackagesBlock/ServicePackagesBlockLayout.svelte";
    import PackageTabList from "@/pages/services/components/PackageTabList/PackageTabList.svelte";
    import SelectedPackage from "@/pages/services/components/SelectedPackage/SelectedPackage.svelte";
    import type {Package} from "@/types/Package.ts";

    type ServicePackagesBlockProps = {
        packages: Package[]
    }

    const {packages}: ServicePackagesBlockProps = $props();
    const tabs = ['Basic', 'Standard', 'Premium'];
    let selectedTab = $state<number>(0);
    const setSelectedTab = (value: number) => selectedTab = value;
</script>

<ServicePackagesBlockLayout>
    <PackageTabList
            setSelectedTab={setSelectedTab}
            selectedTab={selectedTab}
            packages={packages}
            tabs={tabs}
    />
    <SelectedPackage {...packages[selectedTab]}/>
</ServicePackagesBlockLayout>
