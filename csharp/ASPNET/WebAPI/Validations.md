
# Model validation

- Model validation occurs **after model binding** and reports errors where the data, sent from the client, does not meet our validation criteria
-  Both model validation and data binding occur **before** our request reaches an action inside a controller.
- The response status code, when validation fails, should be `422 Unprocessable Entity` - The server understood the content type + syntax of the request, but it was unable to process validation rules
![[Pasted image 20230725211703.png]]

([More built-in validation attributes here](https://docs.microsoft.com/en-us/dotnet/api/system.componentmodel.dataannotations?view=net-5.0))

## Custom Attributes and IValidationObject

```cs
public class ScienceBookAttribute : ValidationAttribute
{
	public BookGenre Genre { get; set; }
	
	public string Error => $"The genre of the book must be {BookGenre.Science}";
	public ScienceBookAttribute(BookGenre genre)
	{
		Genre = genre;
	}
	
	protected override ValidationResult? IsValid(object? value, ValidationContext
	validationContext)
	{
		var book = (Book)validationContext.ObjectInstance;
		if (!book.Genre.Equals(Genre.ToString()))
			return new ValidationResult(Error);
		return ValidationResult.Success;
	}
}
```
How to use it:
```cs
[ScienceBook(BookGenre.Science)]
public string? Genre { get; set; }
```
