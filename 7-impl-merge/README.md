
## A shopping list dataset has the following structure
item_name,is_needed,added_date,quantity

## 1-simple-merge: two shopping lists are given for two different stores. We neeed to merge them based on item_name. Concatenate columns from the RHS using prefix rhs_

LHS:
item_name,is_needed,added_date,quantity

RHS:
item_name,is_needed,added_date,quantity

Result:
item_name,is_needed,added_date,quantity,rhs_is_needed,rhs_added_date,rhs_quantity
