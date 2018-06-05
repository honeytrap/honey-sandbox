import * as React from 'react';
import {
	Badge,
	Card,
	CardBody,
	CardHeader,
	Col,
	Pagination,
	PaginationItem,
	PaginationLink,
	Row,
	Table
} from 'reactstrap';

export function PaginationNav() {
	return (
		<nav>
			<Pagination>
				<PaginationItem><PaginationLink previous
												tag="button">Prev</PaginationLink></PaginationItem>
				<PaginationItem active>
					<PaginationLink
						tag="button">1</PaginationLink>
				</PaginationItem>
				<PaginationItem><PaginationLink
					tag="button">2</PaginationLink></PaginationItem>
				<PaginationItem><PaginationLink
					tag="button">3</PaginationLink></PaginationItem>
				<PaginationItem><PaginationLink
					tag="button">4</PaginationLink></PaginationItem>
				<PaginationItem><PaginationLink next
												tag="button">Next</PaginationLink></PaginationItem>
			</Pagination>
		</nav>
	);
}