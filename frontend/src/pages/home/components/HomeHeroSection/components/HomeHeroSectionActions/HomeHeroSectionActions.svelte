<script>
    import Button from "@/common-components/Button/Button.svelte";
    import HeroSectionActionsLayout
        from "@/pages/home/components/HomeHeroSection/components/HomeHeroSectionActions/HomeHeroSectionActionsLayout.svelte";
    import {Link} from "svelte-routing";
</script>

<HeroSectionActionsLayout>
    <Link to="/search">
        <Button color="white">Find Experts</Button>
    </Link>
    <Link to="/">
        <Button>Create a Service</Button>
    </Link>
</HeroSectionActionsLayout>