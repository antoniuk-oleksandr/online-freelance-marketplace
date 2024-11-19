<script lang="ts">
    import PackageTabLayout from "@/pages/services/components/PackageTab/PackageTabLayout.svelte";

    type PackageTabProps = {
        text: string,
        length: number,
        setSelectedTab: (value: number) => void,
        selectedTab: number,
        index: number,
    }

    const tabProps: PackageTabProps = $props();
    const {text} = tabProps;
</script>

<PackageTabLayout {...tabProps}>
    <span>{text}</span>
</PackageTabLayout>