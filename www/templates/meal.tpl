      <h2>New meal</h2>

      <form role="form" action="/eat" method="post">
        <div class="form-group">
          <label for="typeInput">Type</label> 

		  <select 
          class="form-control" id="typeInput" placeholder=
          "type name" name="typeInput">
			<option>
			Snack
			</option>
			<option>
			Breakfast
			</option>
			<option>
			Lunch
			</option>
			<option>
			Dinner
			</option>
		  </select>

        </div>

        <div class="form-group">
          <label for="productInput">Product</label> <input type=
          "text" class="form-control" id="productInput"
          placeholder="product name" name="productInput">
        </div>

        <div class="form-group">
          <label for="weightInput">Weight (in grams)</label>
          <input type="number" class="form-control" id=
          "weightInput" placeholder="e.g. 100" name="weightInput">
        </div>

        <div class="form-group">
          <label for="proteinInput">Proteins content</label>
          <input type="number" class="form-control" id=
          "proteinInput" placeholder="0" name="proteinInput">
        </div>

        <div class="form-group">
          <label for="fatsInput">Fats content</label> <input type=
          "number" class="form-control" id="fatsInput" placeholder=
          "0" name="fatsInput">
        </div>

        <div class="form-group">
          <label for="carbohydratesInput">Carbohydrates
          content</label> <input type="number" class="form-control"
          id="carbohydratesInput" placeholder="0" name=
          "carbohydratesInput">
        </div>

        <div class="form-group">
          <label for="caloriesInput">calories content</label>
          <input type="number" class="form-control" id=
          "caloriesInput" placeholder="0" name="caloriesInput">
        </div><button type="submit" class=
        "btn btn-primary">Eat!</button>
      </form>

