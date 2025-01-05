<script lang="ts">
    import HeroSectionActions from "@/pages/home/components/HomeHeroSection/components/HomeHeroSectionActions/HomeHeroSectionActions.svelte";
    import HeroSectionFirstLineText from "@/pages/home/components/HomeHeroSection/components/HomeHeroSectionFirstLineText/HomeHeroSectionFirstLineText.svelte";
    import HeroSectionContentLayout
        from "@/pages/home/components/HomeHeroSection/components/HomeHeroSectionContent/HomeHeroSectionContentLayout.svelte";
    import HeroSectionSecondLineText
        from "@/pages/home/components/HomeHeroSection/components/HomeHeroSectionSecondLineText/HomeHeroSectionSecondLineText.svelte";
</script>

<HeroSectionContentLayout>
    <HeroSectionFirstLineText/>
    <HeroSectionSecondLineText/>
    <HeroSectionActions/>
</HeroSectionContentLayout>