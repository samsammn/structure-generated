arch:
	./artisan -entity=$(entity) -module="github.com/Komunitas-Mea/marketplace-mea-api"
arch-full:
	./artisan -entity=$(entity) -type=full -module="github.com/Komunitas-Mea/marketplace-mea-api"
drop:
	rm entity/$(entity).go handler/$(entity).handler.go repository/$(entity).repository.go usecase/$(entity).usecase.go usecase/$(entity).usecase_test.go route/$(entity).route.go