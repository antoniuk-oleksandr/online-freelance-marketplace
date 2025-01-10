<script lang="ts">
    import type {Service} from "@/types/Service.ts";
    import OrderRequestServiceInfoLayout
        from "@/pages/order-request/components/OrderRequestServiceInfo/OrderRequestServiceInfoLayout.svelte";
    import Title from "@/common-components/Title/Title.svelte";
    import OrderRequestServiceInfoImage
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoImage/OrderRequestServiceInfoImage.svelte";
    import OrderRequestServiceInfoContent
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoContent/OrderRequestServiceInfoContent.svelte";
    import OrderRequestServiceInfoTitle
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoTitle/OrderRequestServiceInfoTitle.svelte";
    import OrderRequestServiceInfoRightSide
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoRightSide/OrderRequestServiceInfoRightSide.svelte";
    import OrderRequestServiceInfoBottomLine
        from "@/pages/order-request/components/OrderRequestServiceInfo/components/OrderRequestServiceInfoBottomLine/OrderRequestServiceInfoBottomLine.svelte";

    type OrderRequestServiceInfoProps = {
        serviceData: Service,
    };

    const props: OrderRequestServiceInfoProps = $props();
</script>

<OrderRequestServiceInfoLayout>
    <Title text="Service information"/>
    <OrderRequestServiceInfoContent>
        <OrderRequestServiceInfoImage {...props}/>
        <OrderRequestServiceInfoRightSide>
            <OrderRequestServiceInfoTitle {...props}/>
            <OrderRequestServiceInfoBottomLine {...props}/>
        </OrderRequestServiceInfoRightSide>
    </OrderRequestServiceInfoContent>
</OrderRequestServiceInfoLayout>