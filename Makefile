KINCON_SWAGGER_URL := https://kincone.com/api_docs/swagger.json
KINCONE_SWAGGER := hack/swagger.json

PACKAGE := kinkone

swagger: $(KINCONE_SWAGGER)
$(KINCONE_SWAGGER):
	@wget $(KINCON_SWAGGER_URL) -o $(KINCONE_SWAGGER)

codegen: $(KINCONE_SDK_DIR)
$(KINCONE_SDK_DIR):
	@ swagger generate client -f $(KINCONE_SWAGGER) -A $(PACKAGE) -c $(PACKAGE)
