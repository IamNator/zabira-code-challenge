# zabira-code-challenge


For this particular use case, I was asked to create a sorting solution that is extensible so that different teams can easily add their own sorters without affecting existing code. A flat architecture style could make this more difficult, as any changes to the existing sorting code could potentially affect other parts of the application, making it harder to maintain and extend.

Using an OOP approach, as demonstrated in my implementation, can make the code more modular and easier to extend. By creating an interface for sorters and implementing that interface in different sorter classes, we can add new sorting algorithms without changing the existing code. This approach promotes code reuse and separates concerns, which can lead to a more maintainable and extensible codebase.
