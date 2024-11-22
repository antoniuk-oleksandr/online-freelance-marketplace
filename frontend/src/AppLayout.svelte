<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {useLocation} from "svelte-routing";
    import Header from "@/common-components/Header/Header.svelte";
    import Footer from "@/common-components/Footer/Footer.svelte";
    import PageLayout from "@/common-components/PageLayout.svelte";
    import ToastElement from "@/common-components/ToastElement/ToastElement.svelte";

    let {children}: LayoutProps = $props();

    let location = useLocation();

    const pagesToHide = ["sign-in", "sign-up", "confirm-email", "forgot-password", "reset-password"];
    let hideHeaderFooter = $state<boolean>(pagesToHide.includes($location.pathname.split("/")[1]));

    $effect(() => {
        hideHeaderFooter = pagesToHide.includes($location.pathname.split("/")[1]);
    });
</script>

<div class="flex flex-col gap-y-8 animate-fade-in">
    <ToastElement/>
    {#if !hideHeaderFooter}
        <Header/>
    {/if}
    <PageLayout>
        <div class="min-h-app w-full">
            {@render children()}
        </div>
    </PageLayout>
    {#if !hideHeaderFooter}
        <Footer/>
    {/if}
</div>
