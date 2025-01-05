<script lang="ts">
    import Carousel from "@/common-components/Carousel/Carousel.svelte";
    import HomeProjectsSectionItem
        from "@/pages/home/components/HomeProjectsSection/components/HomeProjectsSectionItem/HomeProjectsSectionItem.svelte";
    import CarouselSlide from "@/common-components/Carousel/components/CarouselSlide/CarouselSlide.svelte";
    import type {BestService} from "@/types/BestService.js";

    type HomeProjectsSectionItemCarouselProps = {
        bestServices: BestService[],
    };

    const {bestServices}: HomeProjectsSectionItemCarouselProps = $props();

    const icons = [
        "hugeicons:code",
        "hugeicons:brush",
        "hugeicons:megaphone-02",
        "hugeicons:smart-phone-01",
    ];
    let services = bestServices.map((service, index) => ({...service, icon: icons[index]}));
</script>

<Carousel slideNumber={services.length}>
    {#each services as project}
        <CarouselSlide>
            <HomeProjectsSectionItem {...project}/>
        </CarouselSlide>
    {/each}
</Carousel>