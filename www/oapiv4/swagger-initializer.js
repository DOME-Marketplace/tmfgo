window.onload = function () {
   //<editor-fold desc="Changeable Configuration Block">

   // the following lines will be replaced by docker/configurator, when it runs in a docker-container
   window.ui = SwaggerUIBundle({
      urls: [
         { url: "/oapiv4/oapiv4/TMF620-productCatalogManagement.json", name: "Product Catalog Management v4" },
         { url: "/oapiv4/oapiv4/TMF622-productOrderingManagement.json", name: "Product Ordering Management v4" },
         { url: "/oapiv4/oapiv4/TMF629-customerManagement.json", name: "Customer Management v4" },
         { url: "/oapiv4/oapiv4/TMF632-party.json", name: "Party Management v4" },
         { url: "/oapiv4/oapiv4/TMF633-serviceCatalogManagement.json", name: "Service Catalog Management v4" },
         { url: "/oapiv4/oapiv4/TMF634-resourceCatalog.json", name: "Resource Catalog Management v4" },
         { url: "/oapiv4/oapiv4/TMF635-usageManagement.json", name: "Usage Management v4" },
         { url: "/oapiv4/oapiv4/TMF637-productInventory.json", name: "Product Inventory Management v4" },
         { url: "/oapiv4/oapiv4/TMF638-serviceInventory.json", name: "Service Inventory Management v4" },
         { url: "/oapiv4/oapiv4/TMF639-resourceInventoryManagement.json", name: "Resource Inventory Management v4" },
         { url: "/oapiv4/oapiv4/TMF641-serviceOrdering.json", name: "Service Ordering Management v4" },
         { url: "/oapiv4/oapiv4/TMF648-quoteManagement.json", name: "Quote Management v4" },
         { url: "/oapiv4/oapiv4/TMF651-agreementManagement.json", name: "Agreement Management v4" },
         { url: "/oapiv4/oapiv4/TMF652-resourceOrderingManagement.json", name: "Resource Ordering Management v4" },
         { url: "/oapiv4/oapiv4/TMF664-resourceFunctionActivation.json", name: "Resource Function Activation v4" },
         { url: "/oapiv4/oapiv4/TMF666-accountManagement.json", name: "Account Management v4" },
         { url: "/oapiv4/oapiv4/TMF669-partyRoleManagement.json", name: "Party Role Management v4" },
         { url: "/oapiv4/oapiv4/TMF678-customerBillManagement.json", name: "Customer Bill Management v4" },
      ],
      dom_id: "#swagger-ui",
      deepLinking: true,
      presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
      plugins: [SwaggerUIBundle.plugins.DownloadUrl],
      layout: "StandaloneLayout",
   });

   //</editor-fold>
};
