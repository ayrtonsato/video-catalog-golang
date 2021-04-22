package services

/*func TestGetCategoriesDbService_GetCategories(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	var fakeCategory = models.Category{
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		DeletedAt:   nil,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should call GetCategories once",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1)
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				_, _ = SUT.GetCategories()
			},
		},
		{
			name: "Should return an slice of models category without errors",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				listCategories := []models.Category{fakeCategory}
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1).
					Return(listCategories, nil)
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				result, err := SUT.GetCategories()
				require.NoError(t, err)
				require.Equal(t, result, listCategories)
			},
		},
		{
			name: "Should throw error if exists",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1).
					Return([]models.Category{}, errors.New("fake_error"))
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				_, err := SUT.GetCategories()
				require.NotEmpty(t, err)
				require.Equal(t, err.Error(), "fake_error")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tc.testCase(t, ctrl)
		})
	}
}

func TestGetCategoriesDbService_GetCategory(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	fakeCategory := models.Category{
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		DeletedAt:   nil,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should get category",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(fakeCategory, nil)
				SUT := NewGetCategoriesDbService(ctgRepository)
				category, err := SUT.GetCategory(uid)
				require.NoError(t, err)
				require.Equal(t, category, fakeCategory)
			},
		},
		{
			name: "Should return ErrNotFound",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, repositories.ErrNoResult)
				SUT := NewGetCategoriesDbService(ctgRepository)
				category, err := SUT.GetCategory(uid)
				require.Error(t, err)
				require.True(t, err.Error() == ErrNotFound.Error())
				require.Equal(t, category, models.Category{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tc.testCase(t, ctrl)
		})
	}
}*/
