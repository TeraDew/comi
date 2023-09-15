// Code generated by ent, DO NOT EDIT.

package book

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBookID holds the string denoting the bookid field in the database.
	FieldBookID = "book_id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldFilePath holds the string denoting the filepath field in the database.
	FieldFilePath = "file_path"
	// FieldBookStorePath holds the string denoting the bookstorepath field in the database.
	FieldBookStorePath = "book_store_path"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldChildBookNum holds the string denoting the childbooknum field in the database.
	FieldChildBookNum = "child_book_num"
	// FieldDepth holds the string denoting the depth field in the database.
	FieldDepth = "depth"
	// FieldParentFolder holds the string denoting the parentfolder field in the database.
	FieldParentFolder = "parent_folder"
	// FieldAllPageNum holds the string denoting the allpagenum field in the database.
	FieldAllPageNum = "all_page_num"
	// FieldFileSize holds the string denoting the filesize field in the database.
	FieldFileSize = "file_size"
	// FieldAuthors holds the string denoting the authors field in the database.
	FieldAuthors = "authors"
	// FieldISBN holds the string denoting the isbn field in the database.
	FieldISBN = "isbn"
	// FieldPress holds the string denoting the press field in the database.
	FieldPress = "press"
	// FieldPublishedAt holds the string denoting the publishedat field in the database.
	FieldPublishedAt = "published_at"
	// FieldExtractPath holds the string denoting the extractpath field in the database.
	FieldExtractPath = "extract_path"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// FieldExtractNum holds the string denoting the extractnum field in the database.
	FieldExtractNum = "extract_num"
	// FieldInitComplete holds the string denoting the initcomplete field in the database.
	FieldInitComplete = "init_complete"
	// FieldReadPercent holds the string denoting the readpercent field in the database.
	FieldReadPercent = "read_percent"
	// FieldNonUTF8Zip holds the string denoting the nonutf8zip field in the database.
	FieldNonUTF8Zip = "non_utf8zip"
	// FieldZipTextEncoding holds the string denoting the ziptextencoding field in the database.
	FieldZipTextEncoding = "zip_text_encoding"
	// EdgePageInfos holds the string denoting the pageinfos edge name in mutations.
	EdgePageInfos = "PageInfos"
	// Table holds the table name of the book in the database.
	Table = "books"
	// PageInfosTable is the table that holds the PageInfos relation/edge.
	PageInfosTable = "single_page_infos"
	// PageInfosInverseTable is the table name for the SinglePageInfo entity.
	// It exists in this package in order to avoid circular dependency with the "singlepageinfo" package.
	PageInfosInverseTable = "single_page_infos"
	// PageInfosColumn is the table column denoting the PageInfos relation/edge.
	PageInfosColumn = "book_page_infos"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBookID,
	FieldOwner,
	FieldFilePath,
	FieldBookStorePath,
	FieldType,
	FieldChildBookNum,
	FieldDepth,
	FieldParentFolder,
	FieldAllPageNum,
	FieldFileSize,
	FieldAuthors,
	FieldISBN,
	FieldPress,
	FieldPublishedAt,
	FieldExtractPath,
	FieldModified,
	FieldExtractNum,
	FieldInitComplete,
	FieldReadPercent,
	FieldNonUTF8Zip,
	FieldZipTextEncoding,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ChildBookNumValidator is a validator for the "ChildBookNum" field. It is called by the builders before save.
	ChildBookNumValidator func(int) error
	// DepthValidator is a validator for the "Depth" field. It is called by the builders before save.
	DepthValidator func(int) error
	// AllPageNumValidator is a validator for the "AllPageNum" field. It is called by the builders before save.
	AllPageNumValidator func(int) error
	// DefaultModified holds the default value on creation for the "Modified" field.
	DefaultModified func() time.Time
)

// OrderOption defines the ordering options for the Book queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBookID orders the results by the BookID field.
func ByBookID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBookID, opts...).ToFunc()
}

// ByOwner orders the results by the Owner field.
func ByOwner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwner, opts...).ToFunc()
}

// ByFilePath orders the results by the FilePath field.
func ByFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFilePath, opts...).ToFunc()
}

// ByBookStorePath orders the results by the BookStorePath field.
func ByBookStorePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBookStorePath, opts...).ToFunc()
}

// ByType orders the results by the Type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByChildBookNum orders the results by the ChildBookNum field.
func ByChildBookNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChildBookNum, opts...).ToFunc()
}

// ByDepth orders the results by the Depth field.
func ByDepth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepth, opts...).ToFunc()
}

// ByParentFolder orders the results by the ParentFolder field.
func ByParentFolder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentFolder, opts...).ToFunc()
}

// ByAllPageNum orders the results by the AllPageNum field.
func ByAllPageNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAllPageNum, opts...).ToFunc()
}

// ByFileSize orders the results by the FileSize field.
func ByFileSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileSize, opts...).ToFunc()
}

// ByAuthors orders the results by the Authors field.
func ByAuthors(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthors, opts...).ToFunc()
}

// ByISBN orders the results by the ISBN field.
func ByISBN(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldISBN, opts...).ToFunc()
}

// ByPress orders the results by the Press field.
func ByPress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPress, opts...).ToFunc()
}

// ByPublishedAt orders the results by the PublishedAt field.
func ByPublishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedAt, opts...).ToFunc()
}

// ByExtractPath orders the results by the ExtractPath field.
func ByExtractPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtractPath, opts...).ToFunc()
}

// ByModified orders the results by the Modified field.
func ByModified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModified, opts...).ToFunc()
}

// ByExtractNum orders the results by the ExtractNum field.
func ByExtractNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtractNum, opts...).ToFunc()
}

// ByInitComplete orders the results by the InitComplete field.
func ByInitComplete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInitComplete, opts...).ToFunc()
}

// ByReadPercent orders the results by the ReadPercent field.
func ByReadPercent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadPercent, opts...).ToFunc()
}

// ByNonUTF8Zip orders the results by the NonUTF8Zip field.
func ByNonUTF8Zip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNonUTF8Zip, opts...).ToFunc()
}

// ByZipTextEncoding orders the results by the ZipTextEncoding field.
func ByZipTextEncoding(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZipTextEncoding, opts...).ToFunc()
}

// ByPageInfosCount orders the results by PageInfos count.
func ByPageInfosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPageInfosStep(), opts...)
	}
}

// ByPageInfos orders the results by PageInfos terms.
func ByPageInfos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPageInfosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPageInfosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PageInfosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PageInfosTable, PageInfosColumn),
	)
}
